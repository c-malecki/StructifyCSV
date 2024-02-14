import { entity } from "../../wailsjs/go/models";

const PropertiesObject = {
  productName: new entity.SchemaProperty({
    type: "string",
    minLength: 3,
    csvHeaderIndex: 0,
  }),
  description: new entity.SchemaProperty({ type: "string", csvHeaderIndex: 1 }),
  featured: new entity.SchemaProperty({ type: "boolean", csvHeaderIndex: 2 }),
  images: new entity.SchemaProperty({
    type: "array",
    items: { type: "string" },
    minItems: 0,
    maxItems: 3,
    csvHeaderIndex: [3, 4],
  }),
  bids: new entity.SchemaProperty({
    type: "object",
    required: ["initialPrice", "lastBid", "totalBids"],
    properties: {
      initialPrice: new entity.SchemaProperty({
        type: "number",
        numMinimum: 5.0,
        numMaximum: 99.99,
        csvHeaderIndex: 5,
      }),
      lastBid: new entity.SchemaProperty({
        type: "null",
        csvHeaderIndex: 6,
      }),
      totalBids: new entity.SchemaProperty({
        type: "integer",
        intMinimum: 0,
        csvHeaderIndex: 7,
      }),
    },
  }),
  seller: new entity.SchemaProperty({
    type: "object",
    properties: {
      id: new entity.SchemaProperty({ type: "integer", csvHeaderIndex: 8 }),
      name: new entity.SchemaProperty({ type: "string", csvHeaderIndex: 9 }),
      address: new entity.SchemaProperty({
        type: "object",
        properties: {
          lineOne: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 10,
          }),
          lineTwo: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 11,
          }),
          city: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 12,
          }),
          state: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 13,
          }),
          country: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 14,
          }),
          postalCode: new entity.SchemaProperty({
            type: "string",
            csvHeaderIndex: 15,
          }),
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
