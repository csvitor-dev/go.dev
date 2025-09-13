import path from "path";
import fs from "fs";

export function getEntriesPattern() {
  const tempParentDir = path.resolve(".tmp");

  if (fs.existsSync(tempParentDir)) {
    fs.rmSync(tempParentDir, { recursive: true, force: true });
  }
  fs.mkdirSync(tempParentDir, { recursive: true });

  const pagesDir = path.resolve("src/pages");
  const filesFromDir = fs
    .readdirSync(pagesDir)
    .filter((file) => file.endsWith(".tsx"));

  return createEntries(filesFromDir, pagesDir, tempParentDir);
}

function createEntries(files: string[], from: string, to: string) {
  const entries: Record<string, string> = {};
  files.forEach((file) => {
    const fileName = path.basename(file, ".tsx");
    const originalPath = path.relative(to, path.resolve(from, fileName));
    const tempFilePath = path.resolve(to, `${fileName}.entry.tsx`);
    const jsxExp = capitalize(fileName);

    const content = `import ReactDOM from "react-dom/client";
import ThemeProvider from "@/components/providers/theme-provider";
import ${jsxExp} from "${originalPath}";

ReactDOM.createRoot(document.getElementById("ui")!).render(<ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <${jsxExp} />
</ThemeProvider>);
`;

    fs.writeFileSync(tempFilePath, content, "utf-8");

    entries[fileName.toLowerCase()] = tempFilePath;
  });

  return entries;
}

function capitalize(arg: string): string {
  return arg
    .split("-")
    .map((s) => s.charAt(0).toUpperCase() + s.substring(1))
    .join("");
}
