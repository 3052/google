# android tv

Alright, here's the full solution for this that worked for me (if anyone else
would need it). A little long and may be a bit complicated, but it works:

1. Find the app on APK mirror and download the .apkm. Even if it doesn't allow
downloading on VMs, this .apkm will work 100% of the time.

2. Create a VM of Android TV 13 (Tiramisu) through android studio, this is the
only version seem to be able to work with this.

3. Install from Google Play Store the APK mirror installer app.

4. When asked to choose a file, simply drag and drop it into the VM, and it
will be in the Downloads folder.

5. Install the app via the APK Mirror Installer app.

6. Download rootavd from here: https://gitlab.com/newbit/rootAVD.

7. Run rootavd, not required to be with admin privileges, and root the device
with it, should be pretty fast and easy.

8. After rootavd succeeds, it will automatically shut down the VM, so start it
again.

9. Open HTTP Toolkit, and click to connect using ADB (the button with the green
android logo in it).

10. ACT FAST HERE: THERE WILL BE A POPUP RELATED TO HTTP TOOLKIT TO GRANT
PERMISSION (will show up for a short period of 1-5 seconds): CLICK GRANT AS
FAST AS YOU CAN, before the HTTP Toolkit app pops up and gets rid of that. That
is important as you can't install right away the CA Certificate to decrypt
HTTPS requests, but if you grant the permission it will work without a problem.

11. After you clicked on the Grant button (if you missed it, just try again.
After a few times, you will be able to press it), the HTTP Toolkit app will say
you need to install the CA Certificate. Click Skip.

12. Restart the machine.

13. You should now be able to connect using HTTP Toolkit with the ADB button
again, and it should show that the CA Certificate is installed (even though we
didn't install anything, it got connected automatically)

14. You can now see all requests of the VM, including HTTPS one's.

Discord user `premie_r`
