import { Button, Flex, PasswordInput, Stack, TextInput, Title } from "@mantine/core";
import { useForm } from '@mantine/form';

export default function SetupForm() {
    const form = useForm({
        mode: 'uncontrolled',
        initialValues: {
            name: '',
            email: '',
            password: '',
        }
    })

    const handleSubmit = async (values: typeof form.values) => {
        const res = await fetch('/v1/admin', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(values),
        })

        if (res.status == 201) {
            window.location.reload()
        }
    }

    return (
        <form onSubmit={form.onSubmit(handleSubmit)}>
            <Flex align="center" justify="space-around" style={{ minHeight: "100vh", padding: "1rem" }}>
                <Stack style={{ width: "100%", maxWidth: 400 }}>
                    <Title size="xl">Setup</Title>

                    <TextInput
                        label="Name"
                        placeholder="Enter your name"
                        required
                        autoFocus
                        disabled={form.submitting}
                        key={form.key('name')}
                        {...form.getInputProps('name')}
                    />

                    <TextInput
                        label="Email"
                        placeholder="Enter your email"
                        required
                        type="email"
                        disabled={form.submitting}
                        key={form.key('email')}
                        {...form.getInputProps('email')}
                    />

                    <PasswordInput
                        label="Password"
                        placeholder="Enter your password"
                        required
                        disabled={form.submitting}
                        key={form.key('password')}
                        {...form.getInputProps('password')}
                    />

                    <Button
                        type="submit"
                        variant="filled"
                        color="dark"
                        disabled={form.submitting}
                        loading={form.submitting}
                        fullWidth>
                        Create Admin Account
                    </Button>
                </Stack>
            </Flex>
        </form>
    )
}