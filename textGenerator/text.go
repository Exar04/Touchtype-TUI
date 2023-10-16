package textgenerator

import (
	"math/rand"
)

var listOfText = []string{
    "In the heart of the bustling city, where time seems to stand still and rush forward simultaneously, people go about their lives, each with a story to tell. In the narrow alleys and busy streets, dreams are born and destinies are shaped. Amidst the cacophony of honking horns and distant sirens, there exists a subtle harmony, an intricate dance of lives intertwining and diverging.",
    "The sun sets, casting a warm golden glow over the urban landscape. Neon signs flicker to life, painting the night with vibrant hues. Street vendors pack their carts, their day's work done, while restaurants and cafes come alive with laughter and the aroma of freshly brewed coffee. The city breathes, its pulse echoing through the skyscrapers that touch the sky, each window holding a glimpse of a different world within.",
    "In the digital realm, countless screens illuminate the darkness. Fingers dance across keyboards, shaping ideas into words and code. Social networks connect people across continents, bridging the gaps between cultures and languages. Information flows like a river, shaping opinions and sparking debates, a testament to the power of human connection in the modern age.",
    "Amidst the technological marvels and the ceaseless urban rhythm, nature weaves its own tapestry. Trees stand tall in city parks, their leaves whispering secrets to the wind. Birds find refuge among the concrete and glass, their songs a reminder of a world beyond the city limits. Even in the heart of this metropolis, the beauty of the natural world finds its place, offering solace to those who seek it.",
    "As night deepens, the city takes on a different character. Quiet descends upon the streets, broken only by the occasional sound of footsteps and distant conversations. The stars above twinkle, their light mingling with the city lights below. It is a moment of reflection, a pause in the endless cycle of activity.",
    "In this block of text, the city comes to life, a living, breathing entity shaped by the dreams and aspirations of its inhabitants. It is a testament to the complexity of human existence, where the old and the new, the natural and the artificial, converge and coexist. Each block, each building, and each person adds a layer to the ever-evolving story of the city, a story that continues to unfold with every passing moment.",
}

func TextGenerator()string{
    return listOfText[rand.Intn(6)]
}
