export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);

    if (url.pathname === "/notion") {
      return await handleNotionRequest(request);
    } else if (url.pathname === "/imgbb") {
      return await handleImgBBRequest(request);
    } else {
      return new Response("Endpoint Not Found", { status: 404 });
    }
  },
};

async function handleNotionRequest(request) {
  const NOTION_TOKEN = "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";
  const DATABASE_ID = "169cc467ec278080abf0f3c9db6d0bbd";

  const body = await request.json();

  const data = {
    parent: {
      database_id: DATABASE_ID,
    },
    properties: {
      "First Name": {
        rich_text: [
          {
            text: {
              content: body.firstName,
            },
          },
        ],
      },
      Patronymic: {
        rich_text: [
          {
            text: {
              content: body.patronymic,
            },
          },
        ],
      },
      "Last Name": {
        rich_text: [
          {
            text: {
              content: body.lastName,
            },
          },
        ],
      },
      Email: {
        rich_text: [
          {
            text: {
              content: body.email,
            },
          },
        ],
      },
      "Phone number": {
        phone_number: body.phoneNumber,
      },
      Country: {
        rich_text: [
          {
            text: {
              content: body.country,
            },
          },
        ],
      },
      "State/Region": {
        rich_text: [
          {
            text: {
              content: body.state,
            },
          },
        ],
      },
      City: {
        rich_text: [
          {
            text: {
              content: body.city,
            },
          },
        ],
      },
      Street: {
        rich_text: [
          {
            text: {
              content: body.street,
            },
          },
        ],
      },
      House: {
        rich_text: [
          {
            text: {
              content: body.house,
            },
          },
        ],
      },
      Apartment: {
        rich_text: [
          {
            text: {
              content: body.apartment,
            },
          },
        ],
      },
      "Zip Code": {
        rich_text: [
          {
            text: {
              content: body.zipcode,
            },
          },
        ],
      },
      Avatar: {
        rich_text: [
          {
            text: {
              content: body.avatar,
            },
          },
        ],
      },
      "Created At": {
        rich_text: [
          {
            text: {
              content: new Date().toISOString(),
            },
          },
        ],
      },
    },
  };

  const response = await fetch("https://api.notion.com/v1/pages", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${NOTION_TOKEN}`,
      "Content-Type": "application/json",
      "Notion-Version": "2022-06-28",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "POST, GET, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type",
    },
    body: JSON.stringify(data),
  });

  const result = await response.json();
  return new Response(JSON.stringify(result), {
    status: 200,
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "POST, GET, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}

async function handleImgBBRequest(request) {
  const IMGBB_API_KEY = "1daa880132b6c031b188ea728f2601a9";

  if (request.method !== "POST") {
    return new Response("Method Not Allowed", { status: 405 });
  }

  const formData = await request.formData();
  const image = formData.get("image"); // The file key must be "image"

  if (!image) {
    return new Response("No image provided", { status: 400 });
  }

  const imgbbResponse = await fetch(
    `https://api.imgbb.com/1/upload?key=${IMGBB_API_KEY}`,
    {
      method: "POST",
      body: formData,
    }
  );

  const result = await imgbbResponse.json();
  return new Response(JSON.stringify(result), {
    status: 200,
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "POST, GET, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}
