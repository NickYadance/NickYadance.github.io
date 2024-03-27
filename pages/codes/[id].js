import {getAllCodeIds, getCodeData} from "../../lib/codes";
import Layout from "../../components/layout";

export default function Code({data}) {
    return (
        <Layout>
            <pre>
                <code className="hljs language-go">
                    <div dangerouslySetInnerHTML={{__html: data.html}}/>
                </code>
            </pre>
        </Layout>
    );
}

export async function getStaticPaths() {
    const paths = getAllCodeIds();
    return {
        paths,
        fallback: false,
    };
}

export async function getStaticProps({params}) {
    const codeData = await getCodeData(params.id);
    return {
        props: {
            data: codeData,
        },
    };
}