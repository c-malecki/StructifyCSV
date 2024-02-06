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
  ["images", "array"],
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
  headerDescriptors: [
    {
      headerText: "Product Name",
      headerIndex: 0,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Product Description",
      headerIndex: 1,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Featured",
      headerIndex: 2,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Image 1",
      headerIndex: 3,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Initial Price",
      headerIndex: 4,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Last Bid",
      headerIndex: 5,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Total Bids",
      headerIndex: 6,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller ID",
      headerIndex: 7,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Name",
      headerIndex: 8,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Address Line 1",
      headerIndex: 9,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Address Line 2",
      headerIndex: 10,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Address City",
      headerIndex: 11,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Address State",
      headerIndex: 12,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Address Country",
      headerIndex: 13,
      schemaProperty: undefined,
      isSelected: false,
    },
    {
      headerText: "Seller Postal Code",
      headerIndex: 14,
      schemaProperty: undefined,
      isSelected: false,
    },
  ],
  usedHeaderIndexes: [],
  // should become reference to file path?
  schema: exampleSchema,
  map: transformForCsvModelMap(exampleSchema.properties),
};
