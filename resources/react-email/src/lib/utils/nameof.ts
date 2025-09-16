export function nameof<T>(name: () => T): [T, string] {
  return [name(), /\w+/.exec(name.toString())![0]];
}
