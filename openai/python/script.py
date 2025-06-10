from openai import OpenAI

client = OpenAI(
    api_key="<KEY>",
    base_url="https://inference-api.genesiscloud.com/openai/v1",
)

response = client.chat.completions.create(
    model="deepseek-ai/DeepSeek-V3-0324",
    messages=[{"role": "user", "content": "Explain to me how AI works"}],
)

print(response.choices[0].message)
