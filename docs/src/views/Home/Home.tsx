import BlocksWrapper from '@/components/BlocksWrapper/BlocksWrapper';
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
                  <td className="method">Get</td>
                  <td className="route">/post/10</td>
                  <td>Get a specific post</td>
                  <td>Yes - Regular user</td>
                </tr>
                <tr>
                  <td className="method" style={{ color: '#ffa500' }}>
                    Patch{' '}
                  </td>
                  <td className="route">/posts/10</td>
                  <td>Update a specific post</td>
                  <td>Yes - Admin only</td>
                </tr>
                <tr>
                  <td className="method" style={{ color: 'green' }}>
                    POST
                  </td>
                  <td className="route">/post</td>
                  <td>Create a new post</td>
                  <td>Yes - Admin only</td>
                </tr>
                <tr>
                  <td className="method" style={{ color: 'red' }}>
                    Delete
                  </td>
                  <td className="route">/delete/10</td>
                  <td>Delete a specific post</td>
                  <td>Yes - Admin only</td>
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
          <BlocksWrapper>
            <h1 className="title">How to Contribute to This Project</h1>

            <p className="paragraph">
              This project is open-source ❤️, I we’d love for you to help make it even better.
              Whether you&apos;re fixing bugs, adding features, or improving documentation, your
              input is always welcome. You can start by visiting the{' '}
              <a
                href="https://github.com/victorlirafront/blog-api"
                target="_blank"
                rel="noopener noreferrer"
                style={{ color: '#007bff', textDecoration: 'none' }}
              >
                GitHub repository
              </a>{' '}
              to explore the code and get involved.
            </p>

            <p className="paragraph">
              Don&apos;t hesitate to fork the project, submit issues, and create pull requests. I
              can&apos;t wait to see what you&apos;ll bring to the project!
            </p>

            <p className="paragraph">
              This API powers this blog:{' '}
              <a href="https://victorlirablog.com.br/" target="_blank" rel="noopener noreferrer">
                https://victorlirablog.com.br
              </a>
            </p>
          </BlocksWrapper>
        </div>
      </div>
    </>
  );
}
