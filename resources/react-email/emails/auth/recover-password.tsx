import {
  Button,
  Container,
  Heading,
  Row,
  Section,
  Text,
} from "@react-email/components";
import BaseLayout from "@/layouts/base-layout";
import { exportProp } from "@/lib/utils";

interface RecoverPasswordProps {
  resetLink: string;
}

export default function RecoverPassword({ resetLink }: RecoverPasswordProps) {
  return (
    <BaseLayout
      previewText="To recover a user's password"
      className="bg-primary"
    >
      <Container className="bg-primary text-primary-foreground p-6">
        <Heading as="h2" className="text-xl font-semibold mb-4">
          Recover password
        </Heading>

        <Section className="my-5">
          <Row>
            <Text className="leading-6 mb-4">
              You have requested a password reset.
            </Text>
            <Text className="leading-6 mb-4">
              Click the button below to continue:
            </Text>
            <Button
              href={exportProp("link", resetLink)}
              className="rounded-lg bg-primary-foreground text-primary px-5 py-2 text-sm font-medium no-underline shadow-sm"
            >
              Reset Password
            </Button>
          </Row>
          <Row>
            <Text className="text-sm text-slate-600 mt-6">
              If you didn't request it, just ignore this email.
            </Text>
          </Row>
        </Section>
      </Container>
    </BaseLayout>
  );
}

RecoverPassword.PreviewProps = {
  resetLink: "https://example.com/",
} as RecoverPasswordProps;
