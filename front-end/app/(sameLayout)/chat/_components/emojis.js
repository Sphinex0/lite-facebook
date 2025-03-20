
export const emojis = [];

export const emojRange = [
    [9728, 9984],      // Miscellaneous Symbols
    [9984, 10176],     // Dingbats
    [127744, 128512],  // Miscellaneous Symbols and Pictographs
    [128512, 128592],  // Emoticons
    [128640, 128704],  // Transport and Map Symbols
    [129280, 129536]   // Supplemental Symbols and Pictographs
];

emojRange.forEach(([i, j]) => {
    for (let index = i; index < j; index++) {
        emojis.push(String.fromCodePoint(index));
    }
})
