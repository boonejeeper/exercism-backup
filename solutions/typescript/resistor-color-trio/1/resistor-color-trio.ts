export function decodedResistorValue(colors: string[]): string {
  const significantDigits = decodedValue(colors.slice(0,2));
  const factor = Math.pow(10,colorCode(colors[2]));
  var prefix = '';


  var value = significantDigits * factor;
  if (value >= 1_000_000_000) {
    value /= 1_000_000_000;
    prefix = 'giga';
  } else if (value >= 1_000_000) {
    value /= 1_000_000;
    prefix = 'mega';
  } else if (value >= 1_000) {
    value /= 1_000;
    prefix = 'kilo';
  }

  return `${value} ${prefix}ohms`;
}

export function decodedValue(colors: string[]): number {
  return (colorCode(colors[0]) * 10) + colorCode(colors[1]) ;
}

export const colorCode = (color: string): number => {
  return COLORS.indexOf(color);
}

export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
];
