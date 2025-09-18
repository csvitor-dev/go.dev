export function nameof<T>(name: () => T) {
  return /\w+/.exec(name.toString())![0];
}
