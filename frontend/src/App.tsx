import {
  createBrowserRouter,
  Outlet,
  RouterProvider
} from 'react-router-dom';
import { Footer } from './components/footer/Footer';
import { Header } from './components/header/Header';
import { Home } from './views/Home';
import { NotFound } from './views/NotFound';

const Layout = (): JSX.Element => {
  return (
    <>
      <Header />
      <main className="bg-gray">
        <Outlet />
      </main>
      <Footer />
    </>
  );
};

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <Home />
      },
      {
        path: '/not-found',
        element: <NotFound />
      }
    ]
  }
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
