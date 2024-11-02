interface IProps {
  onClick?: () => void;
}

export function LogoutButton(props: IProps) {
  return (
    <button
      onClick={props.onClick}
      className="h-[30px] w-[71px] rounded-[3px] border border-white text-sm font-medium uppercase text-white transition-colors duration-300 ease-in-out hover:bg-white/20">
      logout
    </button>
  );
}
