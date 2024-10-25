interface IProps {
  onClick?: () => void;
}

export function SignupButton(props: IProps) {
  return (
    <button
      onClick={props.onClick}
      className="h-[30px] w-[71px] rounded-[3px] bg-white text-sm font-medium uppercase text-[#023f63] transition-colors duration-300 ease-in-out hover:bg-white/80">
      sign up
    </button>
  );
}
