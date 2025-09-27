import type { ComponentProps } from "solid-js"
import FullLogo from "../../../identity/logo-light.svg"
import OrnateLogo from "../../../identity/logo-ornate-light.svg"

export interface LogoProps extends ComponentProps<"svg"> {
  variant?: "mark" | "full" | "ornate"
  size?: number
}

export function Logo(props: LogoProps) {
  const { variant = "mark", size = 64, ...others } = props

  if (variant === "mark") {
    return (
      <svg
        width={size}
        height={size * (42 / 64)}
        viewBox="0 0 64 42"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        class={`text-text ${props.class ?? ""}`}
        {...others}
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M0 0H32V41.5955H0V0ZM24 8.5H8V33H24V8.5Z"
          fill="currentColor"
        />
        <path d="M40 0H64V8.5H48V33H64V41.5H40V0Z" fill="currentColor" />
      </svg>
    )
  }

  if (variant === "full") {
    return <FullLogo width={size * (289 / 42)} height={size} {...others} />
  }

  return <OrnateLogo width={size * (289 / 42)} height={size} {...others} />
}
