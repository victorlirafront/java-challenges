import BlocksWrapper from "@/components/BlocksWrapper/BlocksWrapper";
import Footer from "@/components/Footer/Footer";
import Header from "@/components/Header/Header";
import Head from "next/head";

export default function Home() {
  return (
    <>
      <Head>
        <title>Tech blog API</title>
        <meta name="description" content="Tech blog API" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="https://go.dev/images/favicon-gopher.svg" />
      </Head>
      <div>
        <Header/>
        <div>
          <BlocksWrapper>
            <h1 className="title">What is an API?</h1>
            <p className="paragraph">An API is — in short — a set of dedicated URLs that return pure data responses, in most cases in JSON format — meaning the responses won’t contain the kind of presentational overhead that you would expect in a graphical user interface like a website. I took that great definition from this very understandable article. Please refer to it for more information!</p>
          </BlocksWrapper>
        </div>
        <Footer />
      </div>
    </>
  );
}
