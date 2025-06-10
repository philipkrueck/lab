import OpenAI from "openai";

const client = new OpenAI({
  apiKey: "<KEY>",
  baseURL: "https://inference-api.genesiscloud.com/openai/v1",
});

const response = await client.chat.completions.create({
  model: "deepseek-ai/DeepSeek-V3-0324",
  messages: [
    {
      role: "user",
      content: "explain to me how ai works",
    },
  ],
});

console.log(response.choices?.[0]?.message);
