export function exportGolangProp<T>([value, tag]: [T, string]): string | T {
  return value ?? `{{.${capitalize(tag)}}}`;
}

function capitalize(arg: string) {
  return arg.charAt(0).toUpperCase() + arg.substring(1);
}
