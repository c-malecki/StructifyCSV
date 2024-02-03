import type {
  JsonSchema,
  CsvFile,
  CsvModel,
  PropertiesMapValue,
} from "../types/editor.types";
import { transformForCsvModelMap } from "./transform";

const address = new Map<string, PropertiesMapValue>([
  ["line_one", "string"],
  ["line_two", "string"],
  ["city", "string"],
  ["state", "string"],
  ["country", "string"],
  ["postal_code", "string"],
]);

const bids = new Map<string, PropertiesMapValue>([
  ["initial_price", "number"],
  ["last_bid", "null"],
  ["total_bids", "integer"],
]);

const seller = new Map<string, PropertiesMapValue>([
  ["id", "integer"],
  ["name", "string"],
  ["address", address],
]);

export const product = new Map<string, PropertiesMapValue>([
  ["name", "string"],
  ["description", "string"],
  ["featured", "boolean"],
  // ["images", "array"],
  ["bids", bids],
  ["seller", seller],
]);

export const exampleSchema: JsonSchema = {
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
  properties: product,
};

export const exampleCsvFile: CsvFile = {
  fileName: "Products.csv",
  fileLocation: "/home/meeps/Documents/Products.csv",
};

export const exampleCsvModel: CsvModel = {
  headers: [
    { isSelected: false, header: "Product Name", schemaProperty: null },
    { isSelected: false, header: "Product Description", schemaProperty: null },
    { isSelected: false, header: "Featured", schemaProperty: null },
    // { isSelected: false, header: "Image 1", schemaProperty: null },
    // { isSelected: false, header: "Image 2", schemaProperty: null },
    // { isSelected: false, header: "Image 3", schemaProperty: null },
    { isSelected: false, header: "Initial Price", schemaProperty: null },
    { isSelected: false, header: "Last Bid", schemaProperty: null },
    { isSelected: false, header: "Total Bids", schemaProperty: null },
    { isSelected: false, header: "Seller ID", schemaProperty: null },
    { isSelected: false, header: "Seller Name", schemaProperty: null },
    {
      isSelected: false,
      header: "Seller Address Line 1",
      schemaProperty: null,
    },
    {
      isSelected: false,
      header: "Seller Address Line 2",
      schemaProperty: null,
    },
    { isSelected: false, header: "Seller Address City", schemaProperty: null },
    { isSelected: false, header: "Seller Address State", schemaProperty: null },
    {
      isSelected: false,
      header: "Seller Address Country",
      schemaProperty: null,
    },
    { isSelected: false, header: "Seller Postal Code", schemaProperty: null },
  ],
  usedHeaders: [],
  map: transformForCsvModelMap(exampleSchema.properties),
};
