import com.atlassian.jira.component.ComponentAccessor
import java.text.SimpleDateFormat

// Replace with your actual custom field IDs or names
def storyTypeField = ComponentAccessor.customFieldManager.getCustomFieldObjectByName("Story Type")
def testScopeField = ComponentAccessor.customFieldManager.getCustomFieldObjectByName("Test Scope")
def testingSnapshotField = ComponentAccessor.customFieldManager.getCustomFieldObjectByName("Testing Snapshot")

def validStoryTypes = ["Feature", "Enabler", "NFR"]

// Get field values from the issue
def storyType = issue.getCustomFieldValue(storyTypeField)?.toString()?.trim()
def testScope = issue.getCustomFieldValue(testScopeField)?.toString()?.trim()

// Timestamp formatting
def dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss")
def timestamp = dateFormat.format(new Date())

def snapshotText = ""
snapshotText += "Timestamp Of Initial Transition: ${timestamp}\n"
snapshotText += "Story Type: ${storyType}\n"
snapshotText += "Test Scope: ${testScope}"

if (validStoryTypes.contains(storyType)) {
    issue.setCustomFieldValue(testingSnapshotField, snapshotText)
}
