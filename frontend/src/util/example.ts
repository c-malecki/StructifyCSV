import { entity } from "../../wailsjs/go/models";

const PropertiesObject = {
  productName: new entity.PropertySchema({
    type: "string",
    minLength: 6,
    csvHeaderIndex: 0,
  }),
  description: new entity.PropertySchema({ type: "string", csvHeaderIndex: 1 }),
  featured: new entity.PropertySchema({ type: "boolean", csvHeaderIndex: 2 }),
  images: new entity.PropertySchema({
    type: "array",
    items: { type: "string" },
    minItems: 0,
    maxItems: 3,
    csvHeaderIndex: [3, 4],
  }),
  numArr: new entity.PropertySchema({
    type: "array",
    items: { type: "number" },
    minItems: 2,
    csvHeaderIndex: [16, 17],
  }),
  bids: new entity.PropertySchema({
    type: "object",
    required: ["initialPrice", "lastBid", "totalBids"],
    properties: {
      initialPrice: new entity.PropertySchema({
        type: "number",
        numMinimum: 5.0,
        numMaximum: 99.99,
        csvHeaderIndex: 5,
      }),
      lastBid: new entity.PropertySchema({
        type: "null",
        csvHeaderIndex: 6,
      }),
      totalBids: new entity.PropertySchema({
        type: "integer",
        intMinimum: 0,
        csvHeaderIndex: 7,
      }),
    },
  }),
  seller: new entity.PropertySchema({
    type: "object",
    properties: {
      id: new entity.PropertySchema({ type: "integer", csvHeaderIndex: 8 }),
      name: new entity.PropertySchema({ type: "string", csvHeaderIndex: 9 }),
      address: new entity.PropertySchema({
        type: "object",
        properties: {
          lineOne: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 10,
          }),
          lineTwo: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 11,
          }),
          city: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 12,
          }),
          state: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 13,
          }),
          country: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 14,
          }),
          postalCode: new entity.PropertySchema({
            type: "string",
            csvHeaderIndex: 15,
          }),
        },
      }),
    },
  }),
};

export const exampleSchema = new entity.JsonSchema({
  title: "Product Example",
  description: "Just a test schema to build out the editor UI with.",
  type: "object",
  properties: PropertiesObject,
  required: ["productName", "description", "bids"],
});

export const exampleCsvFile = new entity.CsvFileData({
  fileName: "Products.csv",
  location: "/home/meeps/Documents/Products.csv",
  headers: [
    { column: "A", header: "Product Name" },
    { column: "B", header: "Product Description" },
    { column: "C", header: "Featured" },
    { column: "D", header: "Image 1" },
    { column: "E", header: "Image 2" },
    { column: "F", header: "Initial Price" },
    { column: "G", header: "Last Bid" },
    { column: "H", header: "Total Bids" },
    { column: "I", header: "Seller ID" },
    { column: "J", header: "Seller Name" },
    { column: "K", header: "Seller Address Line 1" },
    { column: "L", header: "Seller Address Line 2" },
    { column: "M", header: "Seller Address City" },
    { column: "N", header: "Seller Address State" },
    { column: "O", header: "Seller Address Country" },
    { column: "P", header: "Seller Postal Code" },
    { column: "Q", header: "Num Array 1" },
    { column: "R", header: "Num Array 2" },
  ],
  referenceRows: [[]],
});
