import BlocksWrapper from '@/components/BlocksWrapper/BlocksWrapper';
import Footer from '@/components/Footer/Footer';
import Header from '@/components/Header/Header';
import Head from 'next/head';

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
        <Header />
        <div>
          <BlocksWrapper>
            <h1 className="title">What is an API?</h1>
            <p className="paragraph">
              An API is — in short — a set of dedicated URLs that return pure data responses, in
              most cases in JSON format — meaning the responses won’t contain the kind of
              presentational overhead that you would expect in a graphical user interface like a
              website. I took that great definition from this very understandable article. Please
              refer to it for more information!
            </p>
          </BlocksWrapper>
          <BlocksWrapper>
            <h1 className="title">Which routes are available?</h1>
            <table className="api-table">
              <thead className="table-header">
                <tr>
                  <th scope="col">Method</th>
                  <th scope="col">Endpoint</th>
                  <th scope="col">Response</th>
                  <th scope="col">Token Required</th>
                </tr>
              </thead>
              <tbody className="table-body">
                <tr>
                  <td className="method">GET</td>
                  <td className="route">/posts</td>
                  <td>Get all blog posts</td>
                  <td>No</td>
                </tr>
                <tr>
                  <td className="method">GET</td>
                  <td className="route">/posts</td>
                  <td>Get all blog posts</td>
                  <td>No</td>
                </tr>
                <tr>
                  <td className="method">GET</td>
                  <td className="route">/posts</td>
                  <td>Get all blog posts</td>
                  <td>No</td>
                </tr>
              </tbody>
            </table>
          </BlocksWrapper>
          <BlocksWrapper>
            <h1 className="title">May I use pagination, sorting and filtering?</h1>
            <p className="paragraph">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur, consectetur?
            </p>
            <h1 className="title">Pagination</h1>
            <table className="api-table">
              <thead className="table-header">
                <tr>
                  <th scope="col">Option</th>
                  <th scope="col">Example</th>
                </tr>
              </thead>
              <tbody className="table-body">
                <tr>
                  <td className="method">limit</td>
                  <td className="route">/character?limit=100</td>
                </tr>
                <tr>
                  <td className="method">page</td>
                  <td className="route">/character?page=2 (limit default is 10)</td>
                </tr>
                <tr>
                  <td className="method">offset</td>
                  <td className="route">/character?offset=3 (limit default is 10)</td>
                </tr>
              </tbody>
            </table>
          </BlocksWrapper>
          <BlocksWrapper>
            <h1 className="title">Sorting</h1>
            <table className="api-table">
              <thead className="table-header">
                <tr>
                  <th scope="col">Examples</th>
                </tr>
              </thead>
              <tbody className="table-body">
                <tr>
                  <td className="route">/character?sort=name:asc</td>
                </tr>
                <tr>
                  <td className="route">/quote?sort=character:desc</td>
                </tr>
              </tbody>
            </table>
          </BlocksWrapper>
        </div>
        <Footer />
      </div>
    </>
  );
}
