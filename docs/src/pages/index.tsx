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
                  <td className="method" style={{color: 'red'}}>Delete</td>
                  <td className="route">/delete/10</td>
                  <td>Delete a specific post</td>
                  <td>Yes</td>
                </tr>
              </tbody>
            </table>
          </BlocksWrapper>
          <BlocksWrapper>
            <h1 className="title">May I use pagination</h1>
            <p className="paragraph">
              Yes, you can add pagination, sorting and filtering options to your API requests.
            </p>
            <h1 className="title">Pagination</h1>
            <table className="api-table">
              <thead className="table-header">
                <tr>
                  <th scope="col">Method</th>
                  <th scope="col">Example</th>
                </tr>
              </thead>
              <tbody className="table-body">
                <tr>
                  <td className="method">GET</td>
                  <td className="route">/posts?page=1&limit=10 (limit default is 10)</td>
                </tr>
                <tr>
                  <td className="method">GET</td>
                  <td className="route">
                    <p>
                      <strong>1st page, 10 items:</strong> <code>/posts?page=1&limit=10</code>
                    </p>
                    <p>
                      <strong>2nd page, 10 items:</strong> <code>/posts?page=2&limit=10</code>
                    </p>
                    <p>
                      <strong>3rd page, 20 items:</strong> <code>/posts?page=3&limit=20</code>
                    </p>
                  </td>
                </tr>
              </tbody>
            </table>
          </BlocksWrapper>
          <BlocksWrapper>
            <h1 className="title">Search</h1>
            <table className="api-table">
              <thead className="table-header">
                <tr>
                <th scope="col">method</th>
                  <th scope="col">Examples</th>
                </tr>
              </thead>
              <tbody className="table-body">
                <tr>
                   <td className="method">GET</td>
                  <td className="route">/search?query=javascript&category=all</td>
                </tr>
                <tr>
                   <td className="method">GET</td>
                  <td className="route">/search?query=typescript&category=web</td>
                </tr>
                <tr>
                   <td className="method">GET</td>
                  <td className="route">/search?query=react&category=mobile</td>
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
