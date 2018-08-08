---
layout: post
title:  "HBase-发现新大陆"
categories: HBase
tags:  Java 大数据 入门
---

* content
{:toc}

![Nick](../img/nick-icon.png)

# HBase-发现新大陆
> [Apache](http://www.apache.org/) HBase™ is the [Hadoop](http://hadoop.apache.org/) database, a distributed, scalable, big data store.  


## 本文目录
 * [HBase Java API示例](#java)
 * [基于HBase的GIS demo](#demo)
 * [附：log4j2 spring pattern](#log4j2)
<h2 id="java" >  HBase Java API示例 </h2>

HBase 2.x 中很多类已经被标注为***Deprecated***，在3.x中将不再使用，网上多数例子属于开发者有责任更新的过时版本:sunglasses:因此自己摸索着写了一些例子并简单测试了一下，供参考。
### 连接的建立
首先建立配置文件，配置目标HBase所使用的zookeeper的地址和端口号。
```
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<configuration>
    <property>
        <name>hbase.zookeeper.quorum</name>
        <value>${Your zookeeper address}</value>
    </property>
    <property>
        <name>hbase.zookeeper.property.clientPort</name>
        <value>${Your zookeeper port}</value>
    </property>

</configuration>
```
用Config类读取配置文件，通常文件放在resources目录下，用相对路径名读取即可。也有说法如果不需要自定义配置，可以不指定配置文件，Config将直接读取HBase的hbase-site.xml。
```
Configuration CONFIG = HBaseConfiguration.create();
CONFIG.addResource(HBASE_CONFIG_FILE);
```
建立连接
```
Connection con = ConnectionFactory.createConnection(CONFIG);
```
官方对Connection的说明：
>Connection creation is a heavy-weight operation. Connection implementations are thread-safe, 
so that the client can create a connection once, and share it with different threads. 
Table and Admin instances, on the other hand, are light-weight and are not thread-safe. 
Typically, a single connection per client application is instantiated and every thread will obtain its own Table instance. 
Caching or pooling of Table and Admin is not recommended.
 
可以看出，Connection适合作为单进程单例对象，线程间共享，因此我把Connection声明为static final，但这样涉及到手动close连接的问题，这里还需要有所斟酌:sleepy::sleepy::sleepy:完整代码如下
```
private static final String HBASE_CONFIG_FILE = "hbase-site.xml";
private static final Configuration CONFIG;
private static Connection CON = null;

static {
	CONFIG  = HBaseConfiguration.create();
	CONFIG.addResource(HBASE_CONFIG_FILE);
	try{
		CON = ConnectionFactory.createConnection(CONFIG);
	} catch (IOException e){
		e.printStackTrace();
	}
}
```
之后在每个操作里都会从这个连接获取表对象。
### 数据操作
* Get
```
/**
 * 单行查询
 * @param tbname 表名
 * @param row rowkey
 */
public static void get(String tbname, String row){
	LOG.info("Get row {} from {}", tbname);
	try (
			Table table = CON.getTable(TableName.valueOf(tbname));
	    ){
		Get get = new Get(row.getBytes());
		Result result = table.get(get);
		LOG.info("Result: {}", result);
	} catch (IOException e) {
		e.printStackTrace();
	}
}
```
* Put

Put的过程与Get类似，只列出部分代码
```
Put put = new Put(rowKey.getBytes())
	.addColumn(columnFamily.getBytes(), qual.getBytes(), val.getBytes());
table.put(put);
```
* Delete
```
Delete delete = new Delete(row.getBytes());
table.delete(delete);
```
* Scan

注意到ResultScanner.class：
>Scans usually do not ship all the matching rows in one RPC to the client,
 but instead do this on a per-row basis. This obviously makes sense
 as rows could be very large and sending thousands, and most likely
 more, of them in one call would use up too many resources, and take a
 long time.
 The ResultScanner converts the scan into a get-like operation, wrapping
 the Result instance for each row into an iterator functionality
   
通常scanner实例保存着服务端资源，很容易占用大量内存，因此操作完成后要及时关闭ResultScanner。部分代码如下：
```
            Scan scan = new Scan()
				.withStartRow(startRow.getBytes())
				.withStopRow(stopRow.getBytes());
			List<Result> results = new ArrayList<>();
			ResultScanner resultScanner = table.getScanner(scan);
			for (Result result : resultScanner){
				LOG.info("Results: {}", result);
			}
			resultScanner.close();
```
### 表操作
HTableDescriptor及相关的（H打头）类已经废弃，替代的是***Builder类，使用起来也很方便，所有表相关的操作都由Admin执行。
* Create

```
/**
 * 创建表，若表已存在则删除
 * @param tbname 目标表名
 */
public static void createTable(String tbname, String[] columnFamily){
	LOG.info("Create table: {}", tbname);
	TableName tableName = TableName.valueOf(tbname);
	try (
			Admin admin = CON.getAdmin();
	    ){
		if (admin.tableExists(tableName)){
			LOG.warn("Table: {} already exists, deleted", tableName);
			admin.disableTable(tableName);
			admin.deleteTable(tableName);
		}
		TableDescriptorBuilder tableDescriptorBuilder = TableDescriptorBuilder.newBuilder(tableName);
		List<ColumnFamilyDescriptor> columnFamilyDescriptors = new ArrayList<>(columnFamily.length);
		for (String family : columnFamily){
			columnFamilyDescriptors.add(
					ColumnFamilyDescriptorBuilder.newBuilder(family.getBytes())
					.build()
					);
		}
		tableDescriptorBuilder.setColumnFamilies(columnFamilyDescriptors);
		admin.createTable(tableDescriptorBuilder.build());
	} catch (IOException e) {
		e.printStackTrace();
	}
}
```
* Drop

与Create类似
```
if (admin.tableExists(tableName)) {
	admin.disableTable(tableName);
	admin.deleteTable(tableName);
}
```
<h2 id="demo" >  基于HBase的GIS Demo</h2>

>来源于<<HBase 实战>>，详细的分析请参阅书籍

问题：K个最近邻居

关键点：
1. 空间上靠近的点，HBase物理存储上也应该靠近
2. Geohash可以把多个值转化成一个唯一值

以Geohash对地点的经度和维度值作hash，得出值作为rowkey，即可把空间信息以一种数据形态存到HBase里。之后对数据库作一定精度
的前缀扫描即可解决问题，代码如下：

```
/**
 * GIS Hbase演示
 * @author shuimen
 */
public class KnnQuery {
	private static final Logger LOG = LoggerFactory.getLogger(KnnQuery.class);
	private final static int PRECISION = 7;
	private static final byte[] TABLE = "wifi".getBytes();
	private static final byte[] FAMILY = "cf".getBytes();
	private static final byte[] ID = "id".getBytes();
	private static final byte[] X_COL = "lon".getBytes();
	private static final byte[] Y_COL = "lat".getBytes();

	private static final String USAGE =
			"KNNQuery lon lat n\n" +
					"  help - print this message and exit.\n" +
					"  lon, lat - query position.\n" +
					"  n - the number of neighbors to return.";

	/**
	 * 固定大小为N的优先队列中插入元素
	 * 维持大根堆序，插入元素比堆顶元素小时poll并入队
	 * @param queue 队列
	 * @param ele 元素
	 * @param n 指定大小
	 */
	private void packQueue(Queue<QueryMatch> queue, QueryMatch ele, int n){
		int size = queue.size();
		if (size == n){
			queue.add(ele);
			queue.poll();
		} else if (size < n){
			queue.add(ele);
		} else {
			LOG.error("Error: queue size exceed");
		}
	}

	/**
	 * K个最近邻居获得
	 * @param queue 存储邻居的队列
	 * @param originHash 目标
	 * @param n 限定大小
	 */
	private void takeKNN(Queue<QueryMatch> queue, GeoHash originHash, int n, Point2D origin){
		List<Result> results = HBaseHandler.gisScan(
				new String(TABLE),
				new String(FAMILY),
				new PrefixFilter(originHash.toBase32().getBytes()));
		if (results == null){
			LOG.error("Error: GIS scan failed");
			return;
		}
		for (Result result : results){
			String hash = new String(result.getRow()),
					id = new String(result.getValue(FAMILY, ID));
			double xCol = Double.parseDouble(
					new String(
							result.getValue(FAMILY, X_COL)));
			double yCol = Double.parseDouble(
					new String(
							result.getValue(FAMILY, Y_COL)));
			Double distance = origin.distance(xCol, yCol);
			QueryMatch queryMatch = new QueryMatch(id, hash, xCol, yCol, distance);
			packQueue(queue, queryMatch, n);
		}
	}
	/**
	 * K个最近邻居(K Nearest Neighbour)查询
	 * @param origin 起点坐标
	 * @param n K
	 */
	public Queue<QueryMatch> queryKNN(Point2D origin, int n){

		Queue<QueryMatch> queue = new PriorityQueue<>(n + 1, Comparator.comparingDouble(QueryMatch::getDistance).reversed());
		GeoHash originHash = GeoHash.withCharacterPrecision(origin.getX(), origin.getY(), PRECISION);
		takeKNN(queue, originHash, n, origin);
		for (GeoHash geoHash : originHash.getAdjacent()){
			takeKNN(queue, geoHash, n, origin);
		}
		return queue;
	}

	public static void main(String[] args) {
		if (args.length != 3){
			LOG.error(USAGE);
			return;
		}
		double lon = Double.parseDouble(args[0]);
		double lat = Double.parseDouble(args[1]);
		int n = Integer.parseInt(args[2]);
		Queue<QueryMatch> queue = new KnnQuery().queryKNN(new Point2D.Double(lon, lat), n);
		Stack<QueryMatch> stack = new Stack<>();
		while (queue.size() != 0){
			stack.push(queue.poll());
		}
		LOG.info("Query completed");
		while (stack.size() != 0){
			LOG.info("Nearest neighbour: {}", stack.pop());
		}
	}
}
```
<h2 id = log4j2>附：log4j2 spring pattern</h2>
```
<PatternLayout
                    disableAnsi="false"
                    pattern="%d{${LOG_DATEFORMAT_PATTERN:-yyyy-MM-dd HH:mm:ss.SSS}} %highlight{${LOG_LEVEL_PATTERN:-%5p}}{FATAL=red blink, ERROR=red, WARN=yellow bold, INFO=green, DEBUG=green bold, TRACE=blue} %style{%5pid}{magenta} [%15.15t] %style{%-40.40C{1.}}{cyan} : %m%n${LOG_EXCEPTION_CONVERSION_WORD:-%xEx}"
            />
```
可输出与spring log相似的格式与颜色。如，前个Demo的查询结果：

![query](../img/gis-output.png)