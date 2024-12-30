export default {
  async fetch(request, env, ctx) {
    try {
      // Отримання параметрів запиту
      const url = new URL(request.url);
      const email = url.searchParams.get("email");
      const name = url.searchParams.get("fullname") || "Friend";
      const listId = url.searchParams.get("listId");

      if (!email || !listId) {
        return new Response(
          "Недостатньо даних: email, name і listId є обов'язковими",
          { status: 400 }
        );
      }

      // Отримання access token
      const tokenResponse = await fetch(
        "https://api.sendpulse.com/oauth/access_token",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            grant_type: "client_credentials",
            client_id: "0d6af2e9f054cceb4ef8858ae625b446", // ваш SendPulse Client ID
            client_secret: "d82df455229d1fb39405bed45900df76", // ваш SendPulse Client Secret
          }),
        }
      );

      const tokenData = await tokenResponse.json();
      const accessToken = tokenData.access_token;

      if (!accessToken) {
        return new Response("Помилка авторизації", { status: 500 });
      }

      // Запит на додавання користувача
      const addUserResponse = await fetch(
        `https://api.sendpulse.com/addressbooks/${listId}/emails`,
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${accessToken}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            emails: [
              {
                email,
                variables: { name },
              },
            ],
          }),
        }
      );

      const sendEmailResponse = await fetch(
        "https://api.sendpulse.com/smtp/emails",
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${accessToken}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: {
              subject: "Laidlon Games",
              from: {
                email: "support@cyberpunkattack.com",
              },
              template: {
                id: 265, // 230,
                variables: {
                  name,
                },
              },
              to: [
                {
                  name,
                  email,
                },
              ],
            },
          }),
        }
      );

      console.log(await sendEmailResponse.json(), new Date());

      return addUserResponse.ok
        ? new Response(
            `Користувач успішно доданий${new Date().toISOString()}`,
            { status: 200 }
          )
        : new Response(
            `Помилка: ${JSON.stringify(await addUserResponse.json())}`,
            { status: 500 }
          );
    } catch (err) {
      return new Response(`Виникла помилка: ${err.message}`, { status: 500 });
    }
  },
};
