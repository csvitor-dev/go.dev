export function exportProp<T>(fieldName: string, value: T): string | T {
  return value ?? `{{.${capitalize(fieldName)}}}`;
}

function capitalize(arg: string) {
  return arg.charAt(0).toUpperCase() + arg.substring(1);
}
