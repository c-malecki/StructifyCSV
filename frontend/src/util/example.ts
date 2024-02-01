const address = new Map<string, string | Map<string, any>>([
  ["line_one", "string"],
  ["line_two", "string"],
  ["city", "string"],
  ["state", "string"],
  ["country", "string"],
  ["postal_code", "string"],
]);

const bids = new Map<string, string | Map<string, any>>([
  ["initial_price", "number"],
  ["last_bid", "null"],
  ["total_bids", "integer"],
]);

const seller = new Map<string, string | Map<string, any>>([
  ["name", "string"],
  ["address", address],
]);

export const product = new Map<string, string | Map<string, any>>([
  ["id", "integer"],
  ["name", "string"],
  ["description", "string"],
  ["featured", "boolean"],
  ["images", "array"],
  ["bids", bids],
  ["seller", seller],
]);

export const exampleSchema = {
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
  properties: product,
};
