import QtQuick 2.0
import Ubuntu.Components 1.1

MainView {
    // objectName for functional testing purposes (autopilot-qt5)
    objectName: "mainView"

    // Note! applicationName needs to match the "name" field of the click manifest
    applicationName: "net-info.rene-so36"

    automaticOrientation: false

    // Removes the old toolbar and enables new features of the new header.
    useDeprecatedToolbar: false

    width: units.gu(100)
    height: units.gu(75)


    Page {
        title: i18n.tr("NetInfo")

        Column {
            id: col
            spacing: units.gu(1)
            anchors {
                margins: units.gu(2)
                fill: parent
            }
            Label {
              id: outLabel
              objectName: "label"
              text: ctrl.message
            }
            Text {
              id: output
              width: parent.width
              height: parent.height / 1.5
              text: ctrl.output
              wrapMode: Text.WordWrap
           }
            Button {
                id: refreshBtn
                onClicked: ctrl.getInterfaces()
                text: "Refresh"
                color: "green"
            }
      }
    }
}

