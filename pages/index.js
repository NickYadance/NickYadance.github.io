import Head from 'next/head';
import Layout, {siteTitle} from '../components/layout';
import utilStyles from '../styles/utils.module.css';
import {getSortedPostsData} from "../lib/posts";
import Link from "next/link";
import Date from '../components/date'
import Image from "next/image";
import email from '../public/images/email.svg'
import github from '../public/images/github.svg'
import {Space, Tooltip} from "antd";

export default function Home({allPostsData}) {
    return (
        <Layout home>
            <Head>
                <title>{siteTitle}</title>
            </Head>
            <section className={utilStyles.headingMd}>
                <p>Programmer & Engineer</p>
                <Space size={"large"}>
                    <Tooltip title="https://nickyadance@gmail.com">
                        <Link href={"https://nickyadance@gmail.com"}><Image src={email} alt="email"/></Link>
                    </Tooltip>
                    <Tooltip title="https://github.com/NickYadance">
                        <Link href={"https://github.com/NickYadance"}><Image src={github} alt="github"/></Link>
                    </Tooltip>
                </Space>
            </section>
            <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
                <ul className={utilStyles.list}>
                    {allPostsData.map(({id, date, title}) => (
                        <li className={utilStyles.listItem} key={id}>
                            <Link href={`/posts/${id}`}>{title}</Link>
                            <br/>
                            <small className={utilStyles.lightText}>
                                <Date dateString={date}/>
                            </small>
                        </li>
                    ))}
                </ul>
            </section>

        </Layout>
    );
}

export async function getStaticProps() {
    const allPostsData = getSortedPostsData();
    return {
        props: {
            allPostsData,
        },
    };
}