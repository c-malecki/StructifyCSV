import type {
  SchemaValues,
  CsvFile,
  CsvModel,
  PropertyMapValue,
} from "../types/editor.types";
import { transformForCsvModelMap } from "./transform";

const address = new Map<string, PropertyMapValue>([
  ["line_one", "string"],
  ["line_two", "string"],
  ["city", "string"],
  ["state", "string"],
  ["country", "string"],
  ["postal_code", "string"],
]);

const bids = new Map<string, PropertyMapValue>([
  ["initial_price", "number"],
  ["last_bid", "null"],
  ["total_bids", "integer"],
]);

const seller = new Map<string, PropertyMapValue>([
  ["id", "integer"],
  ["name", "string"],
  ["address", address],
]);

export const product = new Map<string, PropertyMapValue>([
  ["name", "string"],
  ["description", "string"],
  ["featured", "boolean"],
  ["images", "array"],
  ["bids", bids],
  ["seller", seller],
]);

export const exampleSchema: SchemaValues = {
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
    { isSelected: false, header: "Image 1", schemaProperty: null },
    { isSelected: false, header: "Image 2", schemaProperty: null },
    { isSelected: false, header: "Image 3", schemaProperty: null },
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
  map: transformForCsvModelMap(exampleSchema.properties),
};
