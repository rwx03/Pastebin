import {
  QueryClient,
  QueryClientProvider
} from '@tanstack/react-query';
import {
  createBrowserRouter,
  Outlet,
  RouterProvider
} from 'react-router-dom';
import { Footer } from './components/footer/Footer';
import { Header } from './components/header/Header';
import { Home } from './views/Home';
import { Login } from './views/Login';
import { NotFound } from './views/NotFound';

const queryClient = new QueryClient();

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
      },
      {
        path: '/auth/login',
        element: <Login />
      }
    ]
  }
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  );
}

export default App;
