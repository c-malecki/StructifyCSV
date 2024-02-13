import type { CsvSchemaMap } from "../features/CsvEditor/CsvEditor.types";
import { entity } from "../../wailsjs/go/models";

const PropertiesObject = {
  productName: new entity.Schema({ type: "string", minLength: 3 }),
  description: new entity.Schema({ type: "string" }),
  featured: new entity.Schema({ type: "boolean" }),
  images: new entity.Schema({
    type: "array",
    items: { type: "string" },
    minItems: 0,
    maxItems: 3,
  }),
  bids: new entity.Schema({
    type: "object",
    required: ["initialPrice", "lastBid", "totalBids"],
    properties: {
      initialPrice: new entity.Schema({
        type: "number",
        numMinimum: 5.0,
        numMaximum: 99.99,
      }),
      lastBid: new entity.Schema({
        type: "null",
      }),
      totalBids: new entity.Schema({
        type: "integer",
        intMinimum: 0,
      }),
    },
  }),
  seller: new entity.Schema({
    type: "object",
    properties: {
      id: new entity.Schema({ type: "integer" }),
      name: new entity.Schema({ type: "string" }),
      address: new entity.Schema({
        type: "object",
        properties: {
          lineOne: new entity.Schema({ type: "string" }),
          lineTwo: new entity.Schema({ type: "string" }),
          city: new entity.Schema({ type: "string" }),
          state: new entity.Schema({ type: "string" }),
          country: new entity.Schema({ type: "string" }),
          postalCode: new entity.Schema({ type: "string" }),
        },
      }),
    },
  }),
};

export const exampleSchema = new entity.JsonSchema({
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
  type: "object",
  properties: PropertiesObject,
  required: ["productName", "description", "bids"],
});

export const exampleCsvFile = new entity.CsvFileData({
  fileName: "Products.csv",
  location: "/home/meeps/Documents/Products.csv",
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
});

// export const exampleCsvSchemaMap: CsvSchemaMap = transformForCsvModelMap(
//   exampleSchema.properties
// );
