import { useMutation } from '@tanstack/react-query';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Input } from '../components/ui/input/Input';
import { authService } from '../services/auth.service';

interface IForm {
  email: string;
  password: string;
}

export function Login() {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors }
  } = useForm<IForm>({ mode: 'onSubmit' });

  const { mutate } = useMutation({
    mutationKey: ['login'],
    mutationFn: async (credentials: { email: string; password: string }) => {
      return authService.login(credentials.email, credentials.password);
    },
    onSuccess: () => reset()
  });

  const onSubmit: SubmitHandler<IForm> = (data) => {
    mutate({
      email: data.email,
      password: data.password
    });
    console.log(data);
  };

  return (
    <div className="mx-auto h-screen w-full max-w-[1340px] border border-border bg-gray">
      <div className="mx-2.5 mt-4 flex h-full flex-col items-center pt-2">
        <span className="mb-2 flex w-fit text-xl font-semibold uppercase text-white">
          login page
        </span>
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col items-center">
          <Input
            width="240px"
            {...register('email', {
              required: 'Email is required',
              pattern: {
                value: /^\S+@\S+$/i,
                message: 'Email is invalid'
              }
            })}
            placeholder="Enter email..."
          />
          {errors.email && <span className="text-red-500">{errors.email.message}</span>}
          <Input
            type="password"
            width="240px"
            {...register('password', {
              required: 'Password is required',
              minLength: {
                value: 6,
                message: 'Password must be at least 6 characters'
              }
            })}
            placeholder="Enter password..."
          />
          {errors.password && <span className="text-red-500">{errors.password.message}</span>}
          <button type="submit" className="mt-4 h-10 w-40 bg-zinc-700 font-normal text-white">
            LOGIN
          </button>
        </form>
      </div>
    </div>
  );
}
