<div align="center">
  <a href="https://matterai.so">
    <img
      src="https://matterai.so/favicon.png"
      alt="Matter AI Logo"
      height="64"
    />
  </a>
  <br />
  <p>
    <h3>
      <b>
        Axon Code
      </b>
    </h3>
  </p>
  <p>
    <b>
      AI Coding agent for your Terminal
    </b>
  </p>
  <p>

![Matter Og Image](https://res.cloudinary.com/dxvbskvxm/image/upload/v1759129700/Screenshot_2025-09-29_at_12.35.31_tkxne4.png)

  </p>
</div>

> [!NOTE]  
> This cli agent is a fork of OpenCode cli agent, originally here: [https://github.com/sst/opencode](https://github.com/sst/opencode). Please give them a star for building a 100% OSS version.

---

### Installation

```bash
# YOLO
curl -fsSL https://api.matterai.so/install | bash

# Package managers
npm i -g axoncode-ai@latest        # or bun/pnpm/yarn
brew install sst/tap/axoncode      # macOS and Linux
paru -S axoncode-bin               # Arch Linux
```

> [!TIP]
> Remove versions older than 0.1.x before installing.

#### Installation Directory

The install script respects the following priority order for the installation path:

1. `$axoncode_INSTALL_DIR` - Custom installation directory
2. `$XDG_BIN_DIR` - XDG Base Directory Specification compliant path
3. `$HOME/bin` - Standard user binary directory (if exists or can be created)
4. `$HOME/.axoncode/bin` - Default fallback

```bash
# Examples
axoncode_INSTALL_DIR=/usr/local/bin curl -fsSL https://api.matterai.so/install | bash
XDG_BIN_DIR=$HOME/.local/bin curl -fsSL https://api.matterai.so/install | bash
```

### Documentation

For more info on how to configure axoncode [**head over to our docs**](https://docs.matterai.so).

### Contributing

axoncode is an opinionated tool so any fundamental feature needs to go through a
design process with the core team.

> [!IMPORTANT]
> We do not accept PRs for core features.

However we still merge a ton of PRs - you can contribute:

- Bug fixes
- Improvements to LLM performance
- Support for new providers
- Fixes for env specific quirks
- Missing standard behavior
- Documentation

Take a look at the git history to see what kind of PRs we end up merging.

> [!NOTE]
> If you do not follow the above guidelines we might close your PR.

To run axoncode locally you need.

- Bun
- Golang 1.24.x

And run.

```bash
$ bun install
$ bun dev
```

#### Development Notes

**API Client**: After making changes to the TypeScript API endpoints in `packages/axoncode/src/server/server.ts`, you will need the axoncode team to generate a new stainless sdk for the clients.
