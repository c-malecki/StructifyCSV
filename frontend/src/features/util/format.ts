// import type { Model, Schema, Field } from "../../store/store";

// export const formatFieldJson = (field: Field, acc?: Record<string, any>) => {
//   const display = acc ?? {};
//   const reformatName = field.name.replaceAll(" ", "_").toLowerCase();
//   return Object.assign(display, {
//     [reformatName]: field.dataType.value,
//   });
// };

// export const formatSchemaJson = (schema: Schema, acc?: Record<string, any>) => {
//   const display = {};

//   for (let i = 0; i < schema.fields.length; i++) {
//     const field = schema.fields[i];
//     const fieldObj = formatFieldJson(field);
//     Object.assign(display, fieldObj);
//   }
//   return display;
// };

// export const formatModelJson = (model: Model) => {
//   const display = {};
//   const base = {};
//   const baseSchema = model.schemas[model.baseSchemaIdx];

//   for (let i = 0; i < baseSchema.fields.length; i++) {
//     const field = baseSchema.fields[i];
//     Object.assign(base, formatFieldJson(field));
//   }
//   Object.assign(display, base);
//   for (let i = 0; i < model.schemas.length; i++) {
//     const schema = model.schemas[i];
//     const reformatName = schema.name.replaceAll(" ", "_").toLowerCase();
//     if (model.baseSchemaIdx !== i) {
//       Object.assign(display, { [reformatName]: formatSchemaJson(schema) });
//     }
//   }
//   return display;
// };
