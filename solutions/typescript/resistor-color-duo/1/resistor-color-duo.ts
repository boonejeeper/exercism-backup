export function decodedValue(colors: string[]): number {
  var rValue = 0;
  rValue = (COLORS.indexOf(colors[0]) * 10) + COLORS.indexOf(colors[1]);
  return rValue;
}


export const COLORS = [
  'black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white'
]