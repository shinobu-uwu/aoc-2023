import Bun from "bun";

const file = Bun.file("./input");
const input = await file.text();
let sum = 0;
const digitNames = [
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine",
];

for (const line of input.trim().split("\n")) {
  const digits: any = {};

  for (const name of digitNames) {
    let indices = [];
    let currentIndex = line.indexOf(name);

    while (currentIndex !== -1) {
      indices.push(currentIndex);
      currentIndex = line.indexOf(name, currentIndex + 1);
    }

    for (const index of indices) {
      digits[index] = digitNames.indexOf(name) + 1;
    }
  }

  line.split("").map((c, i) => {
    if (c >= "0" && c <= "9") {
      digits[i] = parseInt(c);
    }
  });

  const keys = Object.keys(digits).map((d) => parseInt(d));
  const min = Math.min(...keys);
  const max = Math.max(...keys);
  sum += digits[min] * 10 + digits[max];
}

console.log(sum);
