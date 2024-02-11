import type { JsonSchema, CsvFile, CsvSchemaMap } from "../types/editor.types";
import type { SchemaProperty } from "../types/properties.types";
import { transformForCsvModelMap } from "./transform";
import {
  StringProperty,
  NumberProperty,
  NullProperty,
  IntegerProperty,
  ObjectProperty,
  BooleanProperty,
  ArrayProperty,
} from "../types/properties.types";

const address = new Map<string, SchemaProperty>([
  ["line_one", new StringProperty()],
  ["line_two", new StringProperty()],
  ["city", new StringProperty()],
  ["state", new StringProperty()],
  ["country", new StringProperty()],
  ["postal_code", new StringProperty()],
]);

const bids = new Map<string, SchemaProperty>([
  ["initial_price", new NumberProperty({ minimum: 5.0 })],
  ["last_bid", new NullProperty()],
  ["total_bids", new IntegerProperty()],
]);

const seller = new Map<string, SchemaProperty>([
  ["id", new IntegerProperty()],
  ["name", new StringProperty()],
  ["address", new ObjectProperty({ properties: address })],
]);

export const product = new Map<string, SchemaProperty>([
  ["name", new StringProperty()],
  ["description", new StringProperty()],
  ["featured", new BooleanProperty()],
  ["images", new ArrayProperty({ items: "string" })],
  ["bids", new ObjectProperty({ properties: bids })],
  ["seller", new ObjectProperty({ properties: seller })],
]);

export const exampleSchema: JsonSchema = {
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
  properties: product,
};

export const exampleCsvFile: CsvFile = {
  fileName: "Products.csv",
  fileLocation: "/home/meeps/Documents/Products.csv",
  headers: [
    "Product Name",
    "Product Description",
    "Featured",
    "Image 1",
    "Image 2",
    "Initial Price",
    "Last Bid",
    "Total Bids",
    "Seller ID",
    "Seller Name",
    "Seller Address Line 1",
    "Seller Address Line 2",
    "Seller Address City",
    "Seller Address State",
    "Seller Address Country",
    "Seller Postal Code",
  ],
};

export const exampleCsvSchemaMap: CsvSchemaMap = transformForCsvModelMap(
  exampleSchema.properties
);
