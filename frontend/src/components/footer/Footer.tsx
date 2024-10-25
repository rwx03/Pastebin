export function Footer() {
  return (
    <div className="bg-gray">
      <footer className="mx-auto h-20 max-w-[1340px] border-t border-t-border flex items-center justify-center text-white/40 font-semibold">
        &copy; {new Date().getFullYear()} All rights
        reserved. Check out my GitHub:&nbsp;
        <a
          href="https://github.com/rwx03"
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-100 hover:underline">
          rwx03
        </a>
      </footer>
    </div>
  );
}
