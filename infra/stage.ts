export const domain = (() => {
  if ($app.stage === "production") return "axoncode.ai"
  if ($app.stage === "dev") return "dev.axoncode.ai"
  return `${$app.stage}.dev.axoncode.ai`
})()

export const zoneID = "430ba34c138cfb5360826c4909f99be8"

// new cloudflare.RegionalHostname("RegionalHostname", {
//   hostname: domain,
//   regionKey: "us",
//   zoneId: zoneID,
// })
