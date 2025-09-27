const stage = process.env.SST_STAGE || "dev"

export default {
  url: stage === "production" ? "https://axoncode.ai" : `https://${stage}.axoncode.ai`,
  console: stage === "production" ? "https://axoncode.ai/auth" : `https://${stage}.axoncode.ai/auth`,
  email: "contact@anoma.ly",
  socialCard: "https://social-cards.sst.dev",
  github: "https://github.com/sst/axoncode",
  discord: "https://axoncode.ai/discord",
  headerLinks: [
    { name: "Home", url: "/" },
    { name: "Docs", url: "/docs/" },
  ],
}
