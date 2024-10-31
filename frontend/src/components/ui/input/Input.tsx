import { forwardRef, InputHTMLAttributes } from 'react';

interface IInput
  extends InputHTMLAttributes<HTMLInputElement> {
  width: string;
}

const Input = forwardRef<HTMLInputElement, IInput>(
  ({ width, ...props }, ref) => {
    return (
      <input
        ref={ref}
        type={props.type || 'text'}
        placeholder={props.placeholder}
        style={{ width: width }}
        className="mb-2 flex h-10 border border-border bg-[#2b2b2b] px-3 py-2 text-[13px] leading-[21px] text-text outline-none"
        {...props}
      />
    );
  }
);

export { Input };
