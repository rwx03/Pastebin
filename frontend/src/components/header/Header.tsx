import { useNavigate } from 'react-router-dom';
import { LoginButton } from '../ui/buttons/LoginButton';
import { PasteButton } from '../ui/buttons/PasteButton';
import { SignupButton } from '../ui/buttons/SignupButton';
import img from './../../assets/images/logo.webp';

export const Header: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div className="h-[50px] border-b border-b-border bg-gray">
      <header className="mx-auto flex h-full max-w-[1340px] items-center justify-between">
        <div className="flex h-full items-center">
          <a className="mr-3 flex h-full items-center text-2xl font-bold text-white" href="/">
            <img className="h-full" src={img} alt="" />
            <span className="uppercase">pastebin</span>
          </a>
          <PasteButton />
        </div>
        <div className="flex gap-x-2">
          <LoginButton onClick={() => navigate('/auth/login')} />
          <SignupButton />
        </div>
      </header>
    </div>
  );
};
