const isNumericInput = (e: KeyboardEvent) => {
  return /^[0-9]$/i.test(e.key);
};

const isModifierKey = (e: KeyboardEvent) => {
  const key = e.key;
  return (
    e.shiftKey === true ||
    key === "End" ||
    key === "Home" ||
    key === "Backspace" ||
    key === "Tab" ||
    key === "Enter" ||
    key === "Delete" ||
    key === "Arrow Left" ||
    key === "Arrow Up" ||
    key === "Arrow Right" ||
    key === "Arrow Down" ||
    ((e.ctrlKey === true || e.metaKey === true) &&
      (key === "A" || key === "C" || key === "V" || key === "X" || key === "Z"))
  );
};

export const enforceNumOnKeyDown = (e: KeyboardEvent) => {
  if (e.key === "-") {
    e.preventDefault();
  }
  if (!isNumericInput(e) && !isModifierKey(e)) {
    e.preventDefault();
  }
};

export const enterNumOnKeyUp = (
  e: KeyboardEvent,
  formValues: Record<any, any>,
  formValueKey: string
) => {
  formValues[formValueKey] = (e.target as HTMLInputElement).value.replace(
    /\D/g,
    ""
  );
};
