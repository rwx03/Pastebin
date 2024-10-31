import { useMutation } from '@tanstack/react-query';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';
import { pasteService } from '../../services/paste.service';
import { CreatePaste } from '../ui/buttons/CreatePaste';

interface IForm {
  message: string;
  title: string;
}

export function PasteForm() {
  const { isAuth } = useAuth();
  const navigate = useNavigate();

  const { register, handleSubmit, reset } = useForm<IForm>();

  const { mutate, isError, isPending, isSuccess } = useMutation({
    mutationKey: ['add paste'],
    mutationFn: async (newPaste: { title: string; content: string }) => {
      return pasteService.addPaste(newPaste.title, newPaste.content);
    },
    onSuccess: () => reset()
  });

  const onSubmit: SubmitHandler<IForm> = (data) => {
    if (!isAuth) {
      navigate('/auth/login');
      return;
    }

    mutate({ title: data.title, content: data.message });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex-col">
      <input
        type="text"
        placeholder="Enter title..."
        className="mb-2 flex h-10 w-8/12 border border-border bg-[#2b2b2b] px-3 py-2 text-[13px] leading-[21px] text-text outline-none"
        {...register('title')}
      />
      <textarea
        placeholder="Enter paste message..."
        className="mb-6 flex h-[300px] min-h-[200px] w-8/12 resize-none overflow-hidden break-words border border-border bg-[#2b2b2b] px-3 py-2 text-[13px] leading-[21px] text-text outline-none"
        {...register('message')}
      />
      <CreatePaste />
      {isError && <span>ERROR</span>}
      {isPending && <span>PENDING</span>}
      {isSuccess && <span>SUCCESS</span>}
    </form>
  );
}
