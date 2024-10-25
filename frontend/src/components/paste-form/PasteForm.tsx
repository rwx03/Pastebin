import { SubmitHandler, useForm } from 'react-hook-form';
import { CreatePaste } from '../ui/buttons/CreatePaste';

interface IForm {
  message: string;
}

export function PasteForm() {
  const { register, handleSubmit } = useForm<IForm>();

  const onSubmit: SubmitHandler<IForm> = (data) =>
    console.log(data);

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="flex-col">
      <textarea
        {...register('message')}
        className="height-[300px] mb-6 flex min-h-[200px] w-8/12 resize-none overflow-hidden break-words border border-border bg-[#2b2b2b] px-3 py-2 text-[13px] leading-[21px] text-text outline-none"
      />
      <CreatePaste />
    </form>
  );
}
