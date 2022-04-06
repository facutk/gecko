`flux-system/gotk-sync`
should include this at the bottom

```
  validation: client
  # Enable decryption
  decryption:
    # Use the sops provider
    provider: sops
    secretRef:
      # Reference the new 'sops-gpg' secret
      name: sops-gpg
```
