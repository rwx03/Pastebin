import { useQuery } from '@tanstack/react-query';
import { useParams } from 'react-router-dom';
import { pasteService } from '../services/paste.service';

export function Paste() {
  const { id } = useParams<{ id: string }>();
  const { data: pasteData } = useQuery({
    queryKey: ['getPaste', id],
    queryFn: async () => {
      return pasteService.getPaste(id!);
    },
    enabled: !!id,
    retry: false,
    staleTime: 1000 * 60 * 5
  });

  return (
    <div className="mx-auto h-screen w-full max-w-[1340px] border border-border bg-gray">
      <section className="flex h-full w-full justify-center">
        <div className="mt-10 flex flex-col items-center gap-y-4 font-semibold text-white max-w-[800px] overflow-y-scroll">
          <h1 className="text-2xl">{pasteData?.title}</h1>
          <h2 className="text-xl">{pasteData?.content}</h2>
        </div>
      </section>
    </div>
  );
}
