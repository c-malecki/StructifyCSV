const hoverColorScheme = {
  1: { hover: "blue-lighten-5", font: "black" },
  2: { hover: "blue-lighten-4", font: "black" },
  3: { hover: "blue-lighten-3", font: "black" },
  4: { hover: "blue-lighten-2", font: "black" },
  5: { hover: "blue-lighten-1", font: "white" },
  6: { hover: "blue-darken-1", font: "white" },
  7: { hover: "blue-darken-2", font: "white" },
  8: { hover: "blue-darken-3", font: "white" },
  9: { hover: "blue-darken-4", font: "white" },
};

export const getHoverColorScheme = (schemaNodeLevel: number) => {
  return hoverColorScheme[schemaNodeLevel as keyof typeof hoverColorScheme];
};

export const getLeftIndent = (schemaNodeLevel: number) => {
  return `margin-left: ${schemaNodeLevel * 16}px`;
};
