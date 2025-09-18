import {
  Body,
  Font,
  Head,
  Html,
  pixelBasedPreset,
  Preview,
  Tailwind,
} from "@react-email/components";

interface BaseLayoutProps {
  className?: string | undefined;
  children: React.ReactNode;
  previewText?: string | undefined;
}

export default function BaseLayout({
  className,
  children,
  previewText,
}: BaseLayoutProps) {
  return (
    <Tailwind
      config={{
        presets: [pixelBasedPreset],
        theme: {
          extend: {
            colors: {
              primary: {
                DEFAULT: "oklch(0.145 0 0)",
                foreground: "oklch(0.985 0 0)",
              },
            },
          },
        },
      }}
    >
      <Html>
        <Head>
          <link rel="preconnect" href="https://fonts.googleapis.com" />
          <link rel="preconnect" href="https://fonts.gstatic.com" />
          <link
            href="https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap"
            rel="stylesheet"
          />
          <Font
            fontFamily="Inter"
            fallbackFontFamily="sans-serif"
            webFont={{
              url: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
              format: "woff2",
            }}
            fontStyle="normal"
            fontWeight="100 400 900"
          />
        </Head>

        {previewText !== undefined && previewText.length > 0 && (
          <Preview>{previewText}</Preview>
        )}

        <Body className={className}>{children}</Body>
      </Html>
    </Tailwind>
  );
}
