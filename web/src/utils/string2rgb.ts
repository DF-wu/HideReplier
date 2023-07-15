/**
 * @see https://gist.github.com/0x263b/2bdd90886c2036a1ad5bcf06d6e6fb37
 */
export function str2rgb(text: string): string {
  let hash = 0;
  if (text.length === 0) return "#000";
  for (let i = 0; i < text.length; i++) {
    hash = text.charCodeAt(i) + ((hash << 5) - hash);
    hash = hash & hash;
  }
  let color = "#";
  for (let i = 0; i < 3; i++) {
    let value = (hash >> (i * 8)) & 255;
    let strVal = "00" + value.toString(16);
    color += strVal.substring(strVal.length - 2);
  }
  return color;
}
