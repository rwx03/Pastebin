import Add from './../../../icons/Add.svg?react';

interface IProps {
  onClick?: () => void;
}

export function PasteButton(props: IProps) {
  return (
    <button
      className="flex h-[1.875rem] items-center rounded-[3px] bg-[#61ba65] pl-1 pr-2.5"
      onClick={props.onClick}>
      <Add />
      <span className="font-semibold text-white">
        paste
      </span>
    </button>
  );
}
