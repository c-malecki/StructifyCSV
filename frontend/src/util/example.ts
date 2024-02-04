import type {
  JsonSchema,
  CsvFile,
  CsvModel,
  PropertiesMapValue,
} from "../types/editor.types";
import { transformForCsvModelMap } from "./transform";
import { entity } from "../../wailsjs/go/models";

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

export const exampleCsvModel = new entity.CsvModel({
  headerDescriptors: [
    {
      isSelected: false,
      headerText: "Product Name",
      headerIndex: 0,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Product Description",
      headerIndex: 1,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Featured",
      headerIndex: 2,
      schemaProperty: undefined,
    },
    // { isSelected: false, headerText: "Image 1", schemaProperty: null },
    // { isSelected: false, headerText: "Image 2", schemaProperty: null },
    // { isSelected: false, headerText: "Image 3", schemaProperty: null },
    {
      isSelected: false,
      headerText: "Initial Price",
      headerIndex: 3,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Last Bid",
      headerIndex: 4,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Total Bids",
      headerIndex: 5,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller ID",
      headerIndex: 6,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Name",
      headerIndex: 7,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Address Line 1",
      headerIndex: 8,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Address Line 2",
      headerIndex: 9,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Address City",
      headerIndex: 10,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Address State",
      headerIndex: 11,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Address Country",
      headerIndex: 12,
      schemaProperty: undefined,
    },
    {
      isSelected: false,
      headerText: "Seller Postal Code",
      headerIndex: 13,
      schemaProperty: undefined,
    },
  ],
  usedHeaderIndexes: [],
  map: transformForCsvModelMap(exampleSchema.properties),
});
