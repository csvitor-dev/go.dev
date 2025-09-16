import { nameof } from "./nameof";

export function golangProp<T>(prop: () => T) {
  return prop() ?? `{{.${capitalize(nameof(prop))}}}`;
}

function capitalize(arg: string) {
  return arg.charAt(0).toUpperCase() + arg.substring(1);
}
