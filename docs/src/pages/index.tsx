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
          <BlocksWrapper />
          <BlocksWrapper />
          <BlocksWrapper />
        </div>
        <Footer />
      </div>
    </>
  );
}
