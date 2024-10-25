import { PasteForm } from '../components/paste-form/PasteForm';

export function Login() {
  return (
    <div className="mx-auto h-screen w-full max-w-[1340px] border border-border bg-gray">
      <div className="mx-2.5 h-full flex-col pt-2">
        <span className="block mb-2 font-semibold text-white">New Paste</span>
        <PasteForm />
      </div>
    </div>
  );
}
