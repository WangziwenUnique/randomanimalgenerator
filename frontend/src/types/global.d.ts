declare global {
  interface Window {
    google?: {
      accounts: {
        id: {
          initialize: (config: {
            client_id: string
            callback: (response: { credential: string }) => void
            auto_select?: boolean
            cancel_on_tap_outside?: boolean
          }) => void
          prompt: () => void
          disableAutoSelect: () => void
          renderButton: (element: HTMLElement | null, options: {
            theme?: string
            size?: string
            type?: string
            shape?: string
            text?: string
            logo_alignment?: string
          }) => void
        }
      }
    }
  }
}

export {}; 